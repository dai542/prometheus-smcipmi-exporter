// -*- coding: utf-8 -*-
//
// © Copyright 2023 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package util

import "log"

type Login struct {
	User     string
	Password string
}

func newLogin(user string, password string) *Login {

	if len(user) == 0 {
		log.Panic("No login user provided")
	}

	if len(password) == 0 {
		log.Panic("No login password provided")
	}

	return &Login{user, password}
}
