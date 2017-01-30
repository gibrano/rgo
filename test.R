options(warn=-1)
args = commandArgs(trailingOnly=TRUE)
eval(parse(text=args[1]))

data = read.table(file, sep=",", header=TRUE)

formula=paste(args[2], "~.", sep="")
model=glm(formula,data,family=binomial(link='logit'))

cat(paste(gsub("\\(|\\)","",names(model$coefficients)), model$coefficients, sep=":"), "/")
