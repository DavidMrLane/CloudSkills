#Create Storage Container
aws s3api create-bucket --bucket $1 --region ${{ env.AWS-REGION}} --create-bucket-configuration LocationConstraint=${{ env.AWS-REGION}}