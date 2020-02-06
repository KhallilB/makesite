# 🔗 makesite

_Create your own custom Static Site Generator (like [Jekyll](https://jekyllrb.com/) or [Hugo](https://gohugo.io/)) by cloning and fulfilling the requirements in this repo!_

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Deliverables](#deliverables)
4. [Resources](#resources)

## Project Structure

```bash
📂 makesite
├── README.md
├── first-post.txt
├── latest-post.txt
├── makesite.go
└── template.tmpl
```

## Getting Started

1. Visit [github.com/new](https://github.com/new) and create a new repository named `makesite`.
2. Run each command line-by-line in your terminal to set up the project:

```bash
$ cd ~/go/src
$ git clone git@github.com:Make-School-Labs/makesite.git
$ cd makesite
$ git remote rm origin
$ git remote add origin git@github.com:YOUR_GITHUB_USERNAME/makesite.git
```

## Deliverables

**For each task**:

- Complete each task in the order they appear.
- Use [GitHub Task List](https://help.github.com/en/github/managing-your-work-on-github/about-task-lists) syntax to update the task list.

### MVP

Complete the MVP as If you finish early, move on to the stretch challenges.

If you get stuck on any step, be sure to print the output to `stdout`!

#### Requirements

- [x] Read in the contents of the provided `first-post.txt` file.
- [x] Edit the provided HTML template (`layout.html`) to display the contents of `first-post.txt`.
- [x] Render the contents of `first-post.txt` using Go Templates.
- [x] Write the HTML template to the filesystem to a file. Name it `first-post.html`.
- [x] Manually test the generated HTML page by double-clicking the `first-post.html` and opening it in your browser.
- [x] Add, commit, and push to GitHub.
- [ ] Add an argument to your CLI command: the name of any `.txt` file in the same directory as your program. Use `latest-post.txt` to test.
- [ ] Update the `save` function to use the input filename to generate a new HTML file. For example, if the input file is named `latest-post.txt`, the generated HTML file should be named `latest-post.html`.
- [ ] Add, commit, and push to GitHub.

#### Stretch Challenges

- [ ] Use Bootstrap, or another CSS framework, to enhance the style and readability of your template. _Get creative! Writing your very own website generator is a great opportunity to broadcast your style, personality, and development preferences to the world!_

## Resources

### Lesson Plans

- [**BEW 2.5**: Project #1 - Static Site Generators](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Ecosystems/#/Lessons/SSGProject)

### Example Code

- [**Go By Example**: Reading Files](https://gobyexample.com/reading-files)
- [**Go By Example**: Writing Files](https://gobyexample.com/writing-files)
- [**Go By Example**: Panic](https://gobyexample.com/panic)
- [**GopherAcademy**: Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [**rapid7.com**: Building a Simple CLI Tool with Golang](https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/)
