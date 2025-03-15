package entity

func (r *Records) LimitRows(limit Limit) {
	*r = (*r)[:limit+1]
}
