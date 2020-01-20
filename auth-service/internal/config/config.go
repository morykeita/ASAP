/**
 * @author Mory Keita on 1/20/20
 */
package config

import "github.com/namsral/flag"

var (
	DataDirectory = flag.String("data-directory","","Path for loading templates and migration scripts")
)


