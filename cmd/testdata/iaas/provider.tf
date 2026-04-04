terraform {
  required_providers {
    outscale = {
      source  = "outscale/outscale"
      version = "1.5.0"
    }
  }
}

provider "outscale" {
}
