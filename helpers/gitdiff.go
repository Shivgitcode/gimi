package helpers

import "os/exec"


func GetGitDiff() (string,error){
	cmd:=exec.Command("git","diff","--cached")
	out,err:=cmd.Output()
	if err!=nil{
		return "",err
	}
	return string(out),nil
}