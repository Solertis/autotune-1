/*
 * Copyright (C) 2017 Acksin, LLC <hi@opszero.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package process

import "github.com/acksin/acksin/stats/memory"

// Process is information about a Linux process
type Process struct {
	// Exe is the executable that is running.
	Exe string
	// PID of the process
	PID int
	// Memory stats of the process
	Memory *memory.ProcessMemory
	// IO contains information about the IO of the machine.
	//	IO *io.ProcessIO
}
