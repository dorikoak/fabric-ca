/*
Copyright IBM Corp. 2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import "os"

// RunMain is the fabric-ca client main
func RunMain(args []string) error {
	// Save the os.Args
	saveOsArgs := os.Args	//cli에서 들어온 매개변수
	os.Args = args

	// Execute the command
	cmdName := ""
	if len(args) > 1 {
		cmdName = args[1]
	}
	ccmd := NewCommand(cmdName)	// ./clientcmd.go의 함수 cmdName을 가진 Command 클래스 생성
	err := ccmd.Execute()	// ccmd.rootCmd.Execute()를 실행시킴, rootCmd는 *cobra.Command 형 필드
	// cobra 타입은 spf13/cobra 패키지에서 선언되어있음
	// cobra 클래스는 cli를 보다 효율적으로 관리할 수 있도록 만들어진 프레임워크,
	// 즉 cli창에서 명령어 입력한는 효과를 낼 수 있음
	// 입력한 명령어를 조합해서 실행시켜줌
 
	// Restore original os.Args
	os.Args = saveOsArgs

	return err
}
