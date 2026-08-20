terraform {
  required_providers {
    outscale = {
      source  = "outscale/outscale"
      version = "1.8.0"
    }
  }
}

provider "outscale" {
}
