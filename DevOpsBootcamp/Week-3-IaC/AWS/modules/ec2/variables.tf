variable "servername" {
    type = string
}
variable "size" {
    type = string
    default = "t2.micro"
}

variable "region" {
    type = string
    default = "eu-west-2"
}