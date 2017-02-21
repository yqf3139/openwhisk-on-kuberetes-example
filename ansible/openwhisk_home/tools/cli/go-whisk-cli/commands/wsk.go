/*
 * Copyright 2015-2016 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
    "github.com/spf13/cobra"
    "../wski18n"
)

// WskCmd defines the entry point for the cli.
var WskCmd = &cobra.Command{
    Use:              "wsk",
    Short:            wski18n.T("OpenWhisk cloud computing command line interface."),
    Long:             logoText(),
    SilenceUsage:     true,
    PersistentPreRunE:parseConfigFlags,
}




func init() {
    WskCmd.SetHelpTemplate(`{{with or .Long .Short }}{{.}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`)

    WskCmd.AddCommand(
        actionCmd,
        activationCmd,
        packageCmd,
        ruleCmd,
        triggerCmd,
        sdkCmd,
        propertyCmd,
        namespaceCmd,
        listCmd,
        apiCmd,
    )

    WskCmd.PersistentFlags().BoolVarP(&flags.global.verbose, "verbose", "v", false, wski18n.T("verbose output"))
    WskCmd.PersistentFlags().BoolVarP(&flags.global.debug, "debug", "d", false, wski18n.T("debug level output"))
    WskCmd.PersistentFlags().StringVarP(&flags.global.auth, "auth", "u", "", wski18n.T("authorization `KEY`"))
    WskCmd.PersistentFlags().StringVar(&flags.global.apihost, "apihost", "", wski18n.T("whisk API `HOST`"))
    WskCmd.PersistentFlags().StringVar(&flags.global.apiversion, "apiversion", "", wski18n.T("whisk API `VERSION`"))
    WskCmd.PersistentFlags().BoolVarP(&flags.global.insecure, "insecure", "i", false, wski18n.T("bypass certificate checking"))
}
