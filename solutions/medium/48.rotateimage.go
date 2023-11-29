// https://leetcode.com/problems/rotate-image/
// 48. Rotate Image
//
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
//
// Example 2:
// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

package medium

func rotate(matrix [][]int) {
	sideLen := len(matrix)

	for r := 0; r < sideLen-1; r++ {
		for c := r + 1; c < sideLen; c++ {
			tmp := matrix[r][c]
			matrix[r][c] = matrix[c][r]
			matrix[c][r] = tmp
		}
	}

	for c := 0; c < sideLen/2; c++ {
		for r := 0; r < sideLen; r++ {
			tmp := matrix[r][c]
			matrix[r][c] = matrix[r][sideLen-1-c]
			matrix[r][sideLen-1-c] = tmp
		}
	}
}

// Another approach
// func rotate(matrix [][]int) {
// 	sideLen := len(matrix)

// 	for r := 0; r < sideLen-1; r++ {
// 		for c := 0; c < sideLen-1-r; c++ {
// 			tmp := matrix[r][c]

// 			matrix[r][c] = matrix[sideLen-1-c][sideLen-1-r]
// 			matrix[sideLen-1-c][sideLen-1-r] = tmp
// 		}
// 	}

// 	for i := 0; i < sideLen/2; i++ {
// 		tmp := matrix[i]
// 		matrix[i] = matrix[sideLen-1-i]
// 		matrix[sideLen-1-i] = tmp
// 	}
// }