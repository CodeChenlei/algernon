// +build !darwin,!dragonfly,!freebsd,!linux,!nacl,!netbsd,!openbsd,!solaris

package main

// Return the server name, version number and description
func banner() string {
	return "\n" + version_string + "\nHTTP/2 web server\n\n"
}
