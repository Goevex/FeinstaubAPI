#!/bin/sh

os_all=(android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos)
archs_all=(386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64)
os_chosen=()
archs_chosen=()

selectContinueBreak="CONTINUE"
selectStartBreak="START"
selectAll="ALL"
selectOsOptions=(${os_all[*]} "$selectAll" "$selectContinueBreak")
selectArchOptions=(${archs_all[*]} "$selectAll" "$selectStartBreak")

echo "Select your os:"
select opt in "${selectOsOptions[@]}"
do
	if [ "$opt" == "$selectAll" ]; then
		os_chosen=(${os_all[*]})
		break
	fi
	if [ "$opt" == "$selectContinueBreak" ]; then
		break
	fi
	os_chosen+=("$opt")
	echo "Your selected os: ${os_chosen[*]}"
done

echo "Select your archs:"
select opt in "${selectArchOptions[@]}"
do
	if [ "$opt" == "$selectAll" ]; then
		archs_chosen=(${archs_all[*]})
		break
	fi
	if [ "$opt" == "$selectStartBreak" ]; then
		break
	fi
	archs_chosen+=("$opt")
	echo "Your selected os: ${archs_chosen[*]}"
done

os_archs=()

for goos in "${os_chosen[@]}"
do
    for goarch in "${archs_chosen[@]}"
    do
		printf "Build ${goos}/${goarch}..."
		mkdir -p ./dev/${goos}/${goarch}
        GOOS=${goos} GOARCH=${goarch} go build main.go 2> ./dev/${goos}/${goarch}/error.log
        if [ $? -eq 0 ]
        then
			echo "success!"
            os_archs+=("${goos}/${goarch}")
			executable="$(find . -maxdepth 1 -type f \( -name "main" -or -name "main.exe" \))"
			if [ "$executable" != "" ]; then
				mv $executable ./dev/${goos}/${goarch}/
			fi
			rm -rf ./dev/${goos}/${goarch}/error.log
		else
			echo "fail"
        fi
    done
done

echo "Successful Archs:"
for os_arch in "${os_archs[@]}"
do
    echo ${os_arch}
done