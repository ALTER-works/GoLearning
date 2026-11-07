package todo

func IdGenerator(tasks []Task) int {
	max_id := 1
	for _, task := range tasks{
		if task.ID >= max_id {
			max_id = task.ID + 1
		}
	}
	return max_id
}