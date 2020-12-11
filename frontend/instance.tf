resource "google_compute_address" "static_ip_cinema" {
  name = "terraform-static-ip-cinema"
  region       = "europe-west3"
  address_type = "EXTERNAL"
}

resource "google_compute_instance" "cinema" {
  name = "cinema"
  machine_type = "e2-standard-2"
  allow_stopping_for_update = "true"

  boot_disk {
    initialize_params {
      image = "ubuntu-1804-lts"
    }
  }

  metadata = {
    ssh-keys = "root:${file(var.gce_ssh_pub_key_file1)} \nroot:${file(var.gce_ssh_pub_key_file2)} \nroot:${file(var.gce_ssh_pub_key_file3)}"
    # ssh-keys = join("\n", [for user, key in var.gce_ssh_pub_key_file : "${user}:${key}"])
  }
  
  network_interface {
    network = "default"
    access_config {
        nat_ip = "${google_compute_address.static_ip_cinema.address}"
    }
  }

  metadata_startup_script="${file("init.sh")}"

    provisioner "file" {
    source      = "./static/"
    destination = "/opt/"
  }

  connection {
    host     = "${google_compute_instance.cinema.network_interface.0.access_config.0.nat_ip}"
    type     = "ssh"
    user     = "${var.gce_ssh_user}"
    # password = "${var.admin_password}"
    private_key = file("./id_rsa")
    agent    = "false"
  }
  # provisioner "local-exec" {
  #     command = "sleep 50; scp -r ./static/** ${var.gce_ssh_user}@${google_compute_instance.cinema.network_interface.0.access_config.0.nat_ip}:/var/www/html/"
  # }  
  
  tags = ["cinema"]
}

resource "google_compute_firewall" "cinema" {
    name = "default-allow-http-terraform"
    network = "default"

    allow {
        protocol = "tcp"
        ports = ["80", "443", "5000"]
    }

    source_ranges = ["0.0.0.0/0"]
    target_tags = ["cinema"]
}

output "ip" {
  value = "${google_compute_instance.cinema.network_interface.0.access_config.0.nat_ip}"
}
