terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.20.0"
    }
    google-beta = {
      version = "4.20.0"
    }
  }
}


locals {
  project_id = "subtle-harmony-349809"
  region     = "us-central1"
}

provider "google" {
  project = local.project_id
  region = local.region
  # Configuration options
}

provider "google-beta" {
  project = local.project_id
  region = local.region
}


resource "google_project_service" "artifact_registry_api" {
  service = "artifactregistry.googleapis.com"

  disable_on_destroy = true
}

resource "google_cloudbuild_trigger" "personal-finance-manager-be-image-trigger" {
  name = "personal-finance-manager-be-image-trigger"
  description = "Trigger used to build image and push it to artifact registry"
  github {
    owner = "nico-iaco"
    name = "personalFinanceManager-be"
    push {
      branch = "main"
    }
  }
  ignored_files = [
    "terraform/**",
    "README.md",
    ".github/**",
    ".gitignore"
  ]

  substitutions = {
    _REPOSITORY = google_artifact_registry_repository.personal-finance-manager-ar.name
    _IMAGE = "personal-finance-manager-be"
  }

  filename = "cloudbuild.yaml"
}

resource "google_artifact_registry_repository" "personal-finance-manager-ar" {
  provider = google-beta

  location = local.region
  repository_id = "personal-finance-manager-ar"
  description = "docker repository for personal finance manager project"
  format = "DOCKER"
  depends_on = [google_project_service.artifact_registry_api]
}

