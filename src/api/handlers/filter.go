package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BuildLPS(pattern string) []int {
    lps := make([]int, len(pattern))
    length := 0
    i := 1

    for i < len(pattern) {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                lps[i] = 0
                i++
            }
        }
    }
	fmt.Println(lps)
    return lps
}

func KMP(ctx *gin.Context) {
	type req struct {
		Text    string `form:"text"`
		Pattern string `form:"pattern"`
	}
	var r req
	err := ctx.ShouldBind(&r)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Build the LPS array
	lps := make([]int, len(r.Pattern))
	j := 0
	for i := 1; i < len(r.Pattern); {
		if r.Pattern[i] == r.Pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	// Perform KMP search
	i, j, counter := 0, 0, 0
	for i < len(r.Text) {
		if r.Pattern[j] == r.Text[i] {
			i++
			j++
		}

		if j == len(r.Pattern) {
			counter++
			j = lps[j-1] // Use LPS array to avoid full reset
		} else if i < len(r.Text) && r.Pattern[j] != r.Text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	ctx.JSON(200, gin.H{
		"number": counter,
	})
}


func BruteForceSearch(ctx *gin.Context) {
	type req struct {
		Text    string `form:"text"`
		Pattern string `form:"pattern"`
	}
	var r req
	err := ctx.ShouldBind(&r)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	textLength := len(r.Text)
	patternLength := len(r.Pattern)
	count := 0

	for i := 0; i <= textLength-patternLength; i++ {
		match := true
		for j := 0; j < patternLength; j++ {
			if r.Text[i+j] != r.Pattern[j] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	ctx.JSON(200, gin.H{
		"number": count,
	})
}
