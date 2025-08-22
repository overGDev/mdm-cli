// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Alvaro Orozco <joaquinorozco2004@gmail.com>
package model

type Section struct {
	Title    string    `yaml:"title"`
	Alias    string    `yaml:"alias,omitempty"`
	Children []Section `yaml:"children,omitempty"`
}
