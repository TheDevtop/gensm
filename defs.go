package main

// Types
type Tape []string
type Table map[string]string
type Dict map[string]Table

// Predefined states
const STATE_MAIN string = "Main"
const STATE_HALT string = "Halt"
const STATE_NULL string = "Null"

// Expression formats
const FMT_BASIC string = "%s : %s => %s\n"
const FMT_SHORT string = "%s => %s\n"
