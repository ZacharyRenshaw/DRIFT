# DRIFT: The Dynamics of Recombination, Inheritance, and Fitness in Time

A comprehensive population genetics modeling software for the creationist community

# Overview

Students of the Bible are often confronted with questions about population growth and population genetics. For example:

- How many people could have been alive at the Flood?
- How many people could have been alive at Babel?
- How could 2.3 million Israelites have left Egypt if only 70 people went in?
- Can you get 2.3 million people in a ‘short’ Egyptian Sojourn of only 215 years?

These have already been addressed in the literature (e.g., [here](https://creation.com/biblical-human-population-growth-model)), but this does not mean every aspect of these questions has been answered. Also, many people might want to explore the question for themselves.

We are sometimes presented with more complex questions that take a specialist to answer. For example:

- How can all people descend from Adam and Eve? Would they have had enough genetic diversity to account for all the different types of people we see today? Would not the inbreeding among their children (and those of Noah) have been bad for humanity?
- Can Swamidass’ [‘Genealogical Adam and Eve’](https://creation.com/review-swamidass-the-genealogical-adam-and-eve) model and Craig’s [‘Adam was a _Homo heidelbergensis_’](https://creation.com/historical-adam-craig) claims be answered by creationists?
- How did the number of differences seen between humans and chimps arise in just a few million years, or is this even possible?
- What are the effects of population bottlenecks like the evolutionary pre-Out of Africa bottleneck or the biblical [Flood bottleneck](https://creation.com/genetic-effects-of-the-flood)?
- Can we explain long-term survival of long-lived species with small populations (e.g., Sanford’s [‘Genetic Entropy’](https://creation.com/mendels-accountant-review))?

To address these questions, and potentially many more, this flexible, open-source population modeling software was created. To make it as accessible to as many people as possible, it is written in the popular programming language Go. Older versions of this software were coded in Perl, C, and Python, but Go was chosen for its speed, flexibility, ease of modification, readability, and ease of incorporating HTML.

# Features

The program can be used to model populations of any size (up to the limits of computer memory) with a large range of possible parameters. The program was designed to reproduce the features of [_Mendel’s Accountant_](https://creation.com/mendels-accountant-review), but the default module uses  overlapping lifespans, instead of _Mendel’s_ discrete generations. Thus, smaller populations can be more easily modeled.

Future modules may include discrete generations (e.g., many plants species) or self-feetlization (e.g., pea plants). Different styles of selection (e.g., birth, annual, probabilty) can also be added. Mutation rates, average mutation effects, etc., are controlled by the user.

One can run models with no mutation effects, so simple population growth experiments are easy to deploy. It is also possible to track the genetic and genealogical contribution of an individual or individuals introduced into the population at any time. ‘Seed’ individual(s) is(are) assigned a digital genome with all bits set to ‘1’. As the generations progress, any individual descended from the seed(s) inherits sections of that person’s digital DNA. A simple recombination model with one recombination per chromosome arm per generation is applied, but this could be modified. One can track the genetic and genealogical descendants of the seed individual(s), the number and average size of recombination blocks, the number of ‘seed’ centromeres remaining in the population, etc. One could also combine mutation with population growth while tracking the descendants of an Adam and an Eve. More advanced users can now answer questions like the maximum number and strength of mutations that a human-like species can withstand, or how much migration between populations is required to completely homogenize them.

The software was designed to be flexible and modifiable. New parameters, subroutines, models, and data calculations can be added easily. In fact, the main program calls six modules, any one of which can be swapped out according to experimental needs: InitializeModel, InitializePop, Birth, Marriage, Death, Save.  

# Requirements

The software is written in the Go programming language. The main program requires Check the first few lines of the main and each module for the list of dependencies. Requirements at present: os, fmt, time, image, image/png, image/color, strings, strconv, math/rand, encoding/csv, and gonum.org/v1/gonum/stat/distuv. In addition, several variable types are defined in drift/types. The modules that need to be included are listed at the beginning of each module and submodule, e.g., "drift/modules/save".

# Installation

To use this software package, Go must be installed on your system. [official Go website](https://go.dev/doc/install). The required libraries listed above can be installed using standard  methods (e.g., run `pip install -r requirements.txt` in the command line).

An installation program (e.g., install.bat) is not included, but setup is simple. Unzip the download file into your directory of choice. If, for example, the user unzipped it into the folder C:/Go/DRIFT, several additional folders will be created:

- c:/Go/drift/modules
- c:/Go/drift/results
- c:/Go/drift/static
- c:/Go/drift/types

This files will appear in the main directory:
- Drift-0.3.go

These files will be in the Data directory:
- actuarial_table.csv, chromosome_data.csv, parameter_defaults.csv,

The Results directory will be empty.

The file system is relational. Thus, the only thing that matters is that the Results and Data subdirectories exist in the folder that contains the program.

# Usage

The method needed to run the program is platform dependent. On a Windows machine, after navigating to the program directory, the command to run the program would be:

     c:\Go\drift> go run drift-0.3.go

**Beware:** If you run multiple models with the same ID, the older data will be overwritten.
**Beware:** Enabling the parameter TrackDead can potentially create very large files. This option is disabled by default. At present, there is no confirmation step when this is enabled.

Enabling Track DNA allows the user to track the DNA and genealogy of a ‘seed’ individual or individuals over time. The seed is added to the population in the year set by the seed year parameter. Currently, the seed is chosen at random. The individual could be male or female and can be of any age. There is no advantage to being the seed (e.g., reduced risk of death or enhanced probability of becoming a parent) and the seed’s descendants are also given no advantages. These are areas that can be easily modified.

Enabling Track Mutations opens a range of additional parameters, including the mutation rate (**mu**), the fraction of neutral mutations [**f(neutral)**], the fraction (of the non-neutral mutations) that are beneficial [**f(beneficial)**], the shape (**Weibull Shape**) and scale (**Weibull Scale**) of the Weibull curve used to set the mutation effects, and an adjustment factor (**Weibull Adjustment**) to reduce the general mutation effects.

Track DNA and Track Mutations use two different engines. Tracking DNA is more memory efficient. When enabled, individuals are assigned two bitstrings **numbits** long. These represent the two copies of the genome in each individual. Numbits is dependent on a genome model that is loaded at the beginning of the run. A default human genome (ChromosomeData.csv) is included in the Data directory. It includes a list of each chromosome arm in the human genome and its length (in megabases). Thus, it is possible to locate any given chromosome arm in a digital genome. These locations are used to control meiosis (explained below).

To save memory, any individual who has zero set bits is deleted from the chromosomes variable.

Tracking mutations is more memory intensive. Any given mutation needs to be assigned both a location and an effect. Every mutation is assigned an ID, effect, posiiotn, dominance, etc. Each individual carries a list of mutations IDs. All mutations in any given bin will either propagate or be lost during meiosis and the fitness effect of any given bin is tabulated by simply summing the effects of the mutations contained in that bin. A histogram of all mutation effects that appear during the model run is stored in memory and saved at the end of the run if the **Mutation Histogram** is enabled in the parameters file.

Currently, the population age distribution is initialized by sampling from an example population (ExamplePop.csv). The age distribution data were generated by using this program to model a static population of 10,000 individuals for 100,000 years. The ages of living people were sampled at the end of the run and saved. In all model runs, survivorship is dictated by an actuarial table (ActuarialTable.csv) that matches the age distribution of an impoverished country obtained from the WHO:[WHO LIFE TABLE FOR 1999: AFR D](who.int/healthinfo/paper09.pdf).

# The Meiosis algorithm

Meiosis is a critical phase in the life cycle of all sexually reproducing organisms, and so it must be represented accurately in these digital organisms. In the current configuration, a random recombination location is chosen for each chromosome arm during the meiosis loop. One of the two chromosome copies is chosen at random and a mask is then generated for the entire genome. For each chromosome, the bits in the mask are then set, either at the center (e.g., from the first recombination point, through the centromere, to the second recombination point) or at the ends, depending on which centromere is chosen. The mask is then applied to the first copy of the individual’s genome with an AND (&) comparison. The inverse (~mask) is applied to the second copy with a second AND comparison. Both copies are then combined with an OR (|) comparison:

     child_copy = (mask & parent_copy_A) | (~mask & parent_copy_B)

This will be the copy of the genome that the child inherits from one parent. The process is simply repeated for the second parent to generate the diploid genome.

Meiosis is used when either Track DNA or Track Mutations are enabled.

# Main Parameters and Settings Frame

## These are the main, user-defined input parameters:

- Model ID: a unique identifier for this model run (change with each run or the data will be overwritten).
- Num Runs: the number of times this model will be repeated. Plots can be saved at the end of each run. All data are saved in the Results directory.
- Start Pop Size: The starting population size.
- Max Pop Size: The maximum population size. Use this for modeling growth or set it equal to Start Pop Size for static populations.
- Max Growth Rate: The maximum population growth rate per year.
- End Year: The number of years to run the model.
- Init Lifespan: The starting lifespan of individuals in the model population. The ‘seed’ individual(s) can have his/her/their own initial lifespan.
- Min Lifespan: The minimum lifespan. This is only used when Init Lifespan is greater than Min Lifespan. Lifespans will drop each generation, but not below this value. There is no ‘Max Lifespan’ because death is controlled by an actuarial table and the death rates of older individuals are quite high. Yet, if an individual is tested for death every year, it is entirely unlikely that any individual could live to ‘biblical’ lifespans, so the probability of death is scaled according to the percent of the maximum lifespan the individual has reached.
- Lifespan Drop: When modeling ‘biblical’ ages, this is the rate at which lifespan decreases per generation. This will bottom out at Min Lifespan.
- Maturity: The age at which males and females can marry.
- Spacing: The minimum number of years between children.
- Birth Prob: The probability (1/x) of an eligible female giving birth in any given year.
- Save Interval: This controls both data plotting and data saving. An extra plot and save are triggered at the end of each run.
- Menopause: The proportion of a female’s maximum lifespan where she stops having children. At present, menopausal woman do not get remarried upon the death of their husband.
- Bottleneck Start: At this specified year, the population will be reduced to Bottleneck Size individuals by randomly killing off individuals.
- Bottleneck End: At this specified year, the population will be allowed to grow, at Max Growth Rate, until it eventually achieves Max Pop Size.
- Bottleneck Size: The size of the population during the bottleneck.
- Track DNA: Activates the DNA Parameters and Settings frame.
- Track Mutations: Activates the Mutations Parameters and Settings frame.
- Track Dead: This will create a file in the Results directory that includes the life history data of every individual born into the population. This allows the user, for example, to create family trees or to assess many other potentially useful statistics. The file size increases linearly with n and runtime (e.g., a population with 1,000 individuals run over 100 years will produce a 2.3 GB file, minimally, but that same population over 1,000 years will create a 26 GB file), so it should be possible to estimate the final size after running a few small prototypes. It should also be possible for an advanced user to programmatically restrict the output data fields to only the ones being studied.
- Max Breeding Inds: This sets the maximum number of adult males and adult non-menopausal females in the population. Excess people will be randomly culled (including children) until this limit is not exceeded. Max Breeding Inds can also be applied to bottlenecks.
- Random Mating: Individuals are assigned a random location within a circle with radius = 0.5 units during the setup loop. Currently, when children are born, they are assigned the latitude and longitude of their father. Two individuals cannot marry if they are located > Random Mating units apart. Set this to ‘1’ for truly random mating.
- Run Model: This will launch the main program. The button will turn red during program execution and return to green when it is finished.
- Seed Year: The year in which the individual(s) whose DNA is to be tracked is introduced.
- Multiplier: To allow for finer recombination, use this to increase the size of the genome. The default size is 3,108 bits, which corresponds to the length of the human genome divided by one million. Chromosome arms range from 153 to 13 bits. This is read from a data file that can easily be modified by the user. Each bit corresponds to one recombination block. More than one mutation can exist in any given recombination block. At present, all mutation effects are additive.
- Initial Heterozygosity: This will set the bits in one copy of each individual’s digital genome to ‘1’, probabilistically, according to the value in this box. If Initial Heterozygosity = 1, every bit in one copy of each individual’s genome will be set. If Initial Heterozygosity = 0.5, one half of the bits in one copy will be set, randomly. Etc.
- Genome Map: This will save a .png file that includes a map of the genome at the top. This is followed by the genomic data for each individual, two lines each.
- All Genome Maps: This will save a unique genome map at each save interval.
- Mu: The mutation rate.
- F(neutral): The proportion of all mutations that are truly neutral.
- F(beneficial): Of the non-neutral mutations, the proportion that are beneficial. For example, if F(neutral) = 0.5 and F(beneficial) = 0.5, beneficial mutations will appear 25% of the time.
- Weibull Shape: One of the two parameters used by the Weibull distribution.
- Weibull Scale: The second parameter used by the Weibull distribution. When Shape = 1 and Scale = 0.5, the Weibull distribution is identical to an exponential distribution.
- Weibull Adjustment: Python has a standard Weibull distribution algorithm, but the values returned (0 to 1) are much too high to be used as mutation effects, so they much be scaled down by 1/x.
- Mutation Histogram: Saves a histogram of the mutation effects of all mutations that ever appeared in the model run and the mutations in circulation at the end of the run. This allows for a quick visual demonstration of the strength of selection.
- Mutation Map: Similar to the DNA map, this creates a .png image with a genome map at the top. Each individual is then represented by two rows. The mutation effect of each genomic bin is represented by the color of the bits in the rows.
- Selection: This is a drop-down with two settings.
  - Annual: Any given individual has a risk of dying each year. The risk is given in an actuarial table loaded at the beginning of the run. When this form of selection is enabled, the risk of dying is increased by the sum of the mutation effects carried by the individual. This is the default setting.
  - Birth: Averages the parent’s mutation burden and subtracts this from birth_probability in the main program. This effectively reduces the chances of high-mutation-burden couples from having children. This is most similar to the way selection is handled in Mendel’s Accountant.

# Program guts

These are three main variables used during a model run:
     model, pop, and mutations
These are custom variables defined in the types file.
- pop.IndData contains life history data for each individual.
- model.FreeParameters is used to track variables that can change during the run (e.g., numinds or max_ID).
- pop.Chromosomes contains two bitarrays per individual, each numbits long. It will stay blank if Track DNA is not selected. To reduce memory,the chromosomes of individuals with zero set bits are deleted. numbits is calculated from the data file ‘chromosome data.csv’ (currently 3046 bits).
- pop.Mutations will stay blank if track mutations is not selected. Otherwise, it will be populated with 2 lists per individual, where each item in the list is, in turn, a list of the mutation IDs they carry at each position.

# Program execution

This is order of operations each model year:

1. Inoculate the population with the seed individuals(s) if year = model.Parameters["inoculation_year"].
2. Birth
3. Marriage
4. Death
8. Save data at specified intervals.

# Example usage

1. The user wants to run a simple population growth model. In the parameters file, they choose a Model ID, set Start Pop Size to 100, and set Max Pop Size to 10,000. They leave everything else at default ad run the model from the command line.
2. The user wants to assess the effects of a population bottleneck. They set the Start Pop Size to 10,000, the Bottleneck Start to year 100, and the Bottleneck End to year 800.
3. The user wishes to know the largest average mutation effect that can be tolerated in a small population. They start by setting Start Pop Size and Max Pop Size to 200. After enablig Track Mutations, they set Mu to 100, f(Neutral) to 0.5, and they enable  Mutation Histogram. After running the model, they adjust the mutation parameters to force the population to survive long-term, changing the Model ID each time. Finally, they open the saved data files in a spreadsheet and graph their results.

# Adding new features
New parameters can easily be added to the file. New modules can be swapped out with the originals also.

# Contributing

We welcome and encourage contributions from the community! If you're interested in enhancing this software, here are several ways you can contribute:

# Reporting Issues

If you come across any bugs, glitches, or unexpected behavior while using the software, please open an issue. When reporting issues, be sure to include detailed information about the problem, including steps to reproduce it and information about your system environment. Your feedback is invaluable in helping us improve the software.

# Feature Requests

Have an idea for a new feature or enhancement? We'd love to hear it! Please submit a feature request on our GitHub repository. Provide a clear description of the proposed feature and its potential benefits to the software. Note that we are depending on the creativity of you, the user, to come up with solutions, hence the reason why we made the program open source!

# Pull Requests

If you have coding skills and would like to contribute directly to the codebase, you can submit a pull request. Follow these steps to contribute:

- Fork the repository: Click the "Fork" button on the top right of the GitHub page.
- Clone your fork: Use git clone to copy the repository to your local machine.
- Create a new branch: Switch to a new branch using git checkout -b feature/your-feature-name.
- Make your changes: Implement the new feature or bug fix in the new branch.
- Test your changes: Ensure that your changes work as expected and don't introduce new issues.
- Commit your changes: Use clear and concise commit messages.
- Push to your fork: Push your changes to your GitHub repository.
- Create a pull request: Submit a pull request from your branch to the main repository's main branch.

Our team will review your pull request, provide feedback, and work with you to integrate the changes into the software.

# Coding Guidelines

To maintain a consistent and readable codebase, please follow standard 'Golang' coding practices and make sure to add explanatory comments wherever practical. By contributing to this software, you become a valuable part of the project's community. We appreciate your efforts and look forward to collaborating with you!

If you have any questions or need assistance, feel free to reach out to us through the issue tracker or contact us via email.

# Acknowledgments

This program was developed mainly by [Dr. Robert Carter](https://creation.com/dr-robert-carter), with help on earlier versions from Chris Hardy and Matthew Powell. Dr. John Sanford’s work on [Genetic Entropy](https://creation.com/mendels-accountant-review) was the main inspiration for this work, and the team that brought that project to fruition cannot be overlooked.

# Contact

For general enquiries, please contact [us@creation.com](us@creation.com).

# Changelog

1 Mar 2024: First public upload of version 1.0.

# Known Issues

No major issues at present, but there are several small bugs.

# Support

This software is free to use. Parties that would like to make a donation in way of thanks can contact Creation Ministries International via [Creation.com](http://www.creation.com/).

# Frequently Asked Questions (FAQ)

1. Do I have to pay to use this program?
   Absolutely not!
2. Can I modify the program?
   Absolutely so!
3. How much memory does it require?
   That depends on the size of the modeled population and whether TrackDNA and TrackMuts are enabled.
4. How much disk space will the output files take up?
   Again, that depends on the model parameters. Enabling TrackDead will save the life history data of every person ever born into the population, which can potentially create very large files.

# Roadmap

1.	GUI:
     a.	GO works well with HTML.
     b.	Launch models from the GUI, allowing users to adjust variables.
     c.	Save variable presets under user-defined model names.
     d.	Add graphing and real-time analysis.
  	e.   Add the ability to save mid-run and launch models from save points.
3.	Optimization
     a.	Parallelization/threading
     b.	Cloud computing
     c.	Algorithmic tweaking
4.	Alternate scenarios:
     a.	Develop a system that allows the user to select among several scenarios (e.g., humans or pea plants).
     b.	The defaults for each scenario will be different.
     c.	A different web form could be developed for each.
5.	New modules:
     a.	The four-alleles test. The challenge has been issued. We can address it.
     b.	ARGWeaver: a more sophisticated four-alleles test using MCMC
     c.	Geography. Mating and migration depend on preloaded maps.
6.	New studies:
     a.	Develop a standard population
     b.	The Genealogical Adam and Eve
     c.	Out of Africa
     d.	Super bottleneck 1MA
     e.	Neanderthals
     f.	Flood/Babel
     g.	Quantifying the effects of natural selection
     h.	Validate Mendel’s Accountant
     i.	Test Jeanson’s ideas about Native American Y chromosome replacement
  	j.   Tribes
7.	Classwork
     a.	With the right GUI and the right module presets and real-time graphs, many possible classroom activities and student projects are possible.


# How to Cite

When reporting the results of any experiment based on this software, please cite:
Carter, R.W., DRIFT: a population genetics modeling software for the creationist community, Creation.com, TBA.

# Disclaimer

We stand by the accuracy of the software algorithms in general, but there is always a possibility that bugs have been inadvertently introduced, so use this at your own discretion. Beware the fact that this software can potentially create very large data files. Finally, any conclusions should be thoroughly doublechecked prior to publication!

# License

This program is released under the [Open Source Initiative MIT license](https://opensource.org/license/mit/). We request that the software not be used for commercial purposes and that proper attribution (see How to Cite above) is always given.
