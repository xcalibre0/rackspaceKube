provider "google" 
{
  # If you run from an instance inside gce with approriate permissions
  # then you dont need to supply a key

  project = "vertical-tuner-193123"
  region = "us-central1"
}

resource "google_container_cluster" "cluster" {
  name = "key-value-thingy"
  zone = "us-central1-a"

  initial_node_count = 3
  
  node_config {
    tags = [ "key-value-node" ]

    machine_type = "g1-small"

    disk_size_gb = 10
   
    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

  }

}
