#Create Storage Container
if aws s3api bucket-not-exists --bucket $1
then aws s3api create-bucket --bucket $1 --region $2 --create-bucket-configuration LocationConstraint=$2
fi