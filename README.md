<h1>Project</h1>

This is a simple prototype using Greptile API to search for potentially inefficient code in a codebase using predefined prompts.


It currently returns a list of files that may contain inefficient code. For future improvements, it could return the function and line numbers that the API has identified.

<h2>Idea</h2>

The idea behind this prototype is to build a tool that will statically analyze the performance of a codebase and identify areas for potential improvements.

It could be used as a starting point for identifying inefficient code. However, further analysis would be needed. It's important to profile the code with realistic use cases and have an understanding of what areas of the code needs to be most responsive.


<h2>How to run this</h2>

1. Set environment variables for `GREPTILE_API_KEY` and `GITHUB_TOKEN`. 
    - See [Greptile documentation](https://docs.greptile.com/quickstart#permissions) for more details.
2. Update the `repository` and `branch` in `search()` in `greptile_project.go`.
    - This step assumes that the repository has already been [indexed](https://docs.greptile.com/api-reference/index)
3. Run `go build .` in the project directory.
4. Run `greptile_project.exe`.