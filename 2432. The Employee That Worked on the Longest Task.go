// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/description/

func hardestWorker(n int, logs [][]int) int {
    win_id, win_time := logs[0][0], logs[0][1]
    for i := 1; i < len(logs); i++ {
        cur_id := logs[i][0]
        cur_time := logs[i][1] - logs[i-1][1]
        if cur_time > win_time {
            win_time = cur_time
            win_id = cur_id
        } else if cur_time == win_time && cur_id < win_id {
            win_id = cur_id
        }
    }
    return win_id
}

