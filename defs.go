package main

type Tape []string
type Table map[string]string
type Dict map[string]Table

const STATE_MAIN string = "Main"
const STATE_HALT string = "Halt"
const STATE_NULL string = "Null"

const FMT_BASIC string = "%s : %s => %s\n"
const FMT_ERROR string = "%s => %s\n"
