let students=[{subject:'math',score:90},{subject:'phy',score:98},{subject:'computer science',score:95}]
let totalscore=0;
for(let{score} of students){
    totalscore+=score;
}
console.log(totalscore)