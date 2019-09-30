package problem0735

func asteroidCollision(asteroids []int) []int {
    answers := make([]int, 0)
    
    for _, v := range asteroids{
        keep := true
        for len(answers) > 0 && v < 0 && answers[len(answers)-1] > 0{
            if answers[len(answers)-1] < -v{
                answers = answers[:len(answers)-1]
                continue
            }else if answers[len(answers)-1] == -v{
                answers = answers[:len(answers)-1]
            }
            keep = false
            break            
        }
        
        if keep{
            answers = append(answers, v)
        }   
    }
    return answers
}

