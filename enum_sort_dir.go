package betting

var ESortDir = eSortDir{
	EARLIEST_TO_LATEST: "EARLIEST_TO_LATEST",
	LATEST_TO_EARLIEST: "LATEST_TO_EARLIEST",
}

type eSortDir struct {
	EARLIEST_TO_LATEST eSortDirInternal
	LATEST_TO_EARLIEST eSortDirInternal
}

type eSortDirInternal string
