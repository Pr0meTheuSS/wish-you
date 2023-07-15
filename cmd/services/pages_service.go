package service

import (
	"os"
	"path"
	"strings"
)

/*
 * Project: I-wish-you
 * Created Date: Saturday, July 15th 2023, 11:18:46 am
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

const (
	pagesDirPath   = "cmd/pages/"
	pagesExtension = ".html"
)

func getPagesDirectoryName(pageName string) string {
	return strings.Split(pageName, "-")[0] + "-pages/"
}

func LoadPage(pageName string) ([]byte, error) {
	return os.ReadFile(path.Join(pagesDirPath, getPagesDirectoryName(pageName), pageName+pagesExtension))
}
