package task

type TaskCreate struct {
    Name string `json: "name"`
    Detail string `json: "detail"`
}

type TaskUpdate struct {
    Name *string `json: "name"`
    Detail *string `json: "detail"`
}
