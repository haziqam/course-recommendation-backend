package utils

type IScanner interface {
	Scan(dest ...interface{}) error
}
