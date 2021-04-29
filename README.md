# CloudSkills

This is my repo for working on courses, labs, projects from Cloudskills.io Cloud Native DevOps Bootcamp

I will only be doing the AWS labs to begin with as Azure only give a 30 day free trial, whereas AWS is 1 year!

## Week 5 - CI/CD
As I'm not doing the azure labs yet, I decided to try and deploy my AWS terraform code using gh-actions. I setup TerraformCloud to do the runs as it has built in state storage, to make it work I had to add the tfcloud backend provider in to the main.tf from Week 3.
