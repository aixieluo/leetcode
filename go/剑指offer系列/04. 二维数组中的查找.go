package main

func main()  {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for y:=0; y < len(matrix); y++ {
		for x:=0; x< len(matrix[0]); x++ {
			if matrix[y][x] == target {
				return true
			}
		}
	}
	return false
}
