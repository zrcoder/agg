package hanoi

import "github.com/zrcoder/agg/pkg"

const (
	maxDiskCount = 6

	easyLevelInfo = `Welcome to the Tower of Hanoi!

Goal: Move all disks from pile A to pile C, following these rules:
1. Move only one disk at a time
2. Take the top disk from a pile and place it on another pile or an empty pile
3. Never place a larger disk on top of a smaller disk
`
	easyLevelCode = `// Tip: Type the letter (A, B, or C) to 'select' a pile

// Finish the solution:
// Move the small disk from A to B
A
B
// Move the big disk from A to C
A
C
// Move the small disk from B to C

`

	mediumLevelInfo = `There are more disks now, try your best.
`
	mediumLevelCode = `// You can also type 'a', 'b', or 'c' instead of 'A', 'B', or 'C'.
`

	hardLevelInfo = `Solving the Tower of Hanoi with Recursion
Challenge: 

Implement the core logic of moving the largest disk
Key insights:
1. Recursion breaks down complex problem into simpler sub-problems
2. Each recursive call reduces the problem size by 1 disk
`

	hardLevelCode = `func solve(disks int, a, b, c Pile) {
    // Base case: No disks to move
    if disks == 0 {
        return
    }

    // Step 1: Move n-1 disks from source to auxiliary pile
    // This creates space to move the largest disk
    solve disks-1, a, c, b
    
    // Step 2: Move the largest disk from source to destination
    // TODO: Implement disk movement logic
    
    
	
    // Step 3: Move n-1 disks from auxiliary pile to destination
    // Completes the recursive solution
    solve disks-1, b, a, c
}

// Demonstrate solving Tower of Hanoi with 5 disks
solve 5, a, b, c
`

	expertLevelInfo = "Solve it with recursion"
	expertLevelCode = `func solve(disks int, a, b, c Pile) {

}
`
)

var levels = []pkg.Level{
	{Name: "Easy", Value: 2, Help: pkg.Help{Title: "Level: Easy", Info: easyLevelInfo, Code: easyLevelCode}},
	{Name: "Medium", Value: 3, Help: pkg.Help{Title: "Level: Medium", Info: mediumLevelInfo, Code: mediumLevelCode}},
	{Name: "Hard", Value: 5, Help: pkg.Help{Title: "Level: Hard", Info: hardLevelInfo, Code: hardLevelCode}},
	{Name: "Expert", Value: 6, Help: pkg.Help{Title: "Level: Expert", Info: expertLevelInfo, Code: expertLevelCode}},
}
