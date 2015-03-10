package stats

type metric interface {
	CollectedFor() int // person_id !?
	//person Person // identify who this metric is for !?
	//UUID string // identify who this metric is for !?

	//trx_id string // identify the transaction that this is associated with (so we can find other associated metrics)
	Context() string // transaction id
	Key() string
	Value() int // perhaps other metrics are intergers ?! (i.e. floats)

	CollectedAt() int // epoch-timestamp must have!!!
	// a persistence strategy function
	// shouldn't cloud-up this data-structure...

	// methods to persist them (to JSON !?)
	// or to a DB
}
