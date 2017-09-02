package panel


func CopyFile(in, out string, head bool) {
	Load(in, head).Unload(out)
}
