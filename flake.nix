{
  description = "Go API (grups) dev environment with Postgres";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.git
            pkgs.sqlite
            pkgs.gnumake
          ];

          shellHook = ''
            export CGO_ENABLED=0
            export GOFLAGS='-mod=mod'
            echo "Entered dev shell. Run migrations: make migrate"
            echo "Then run: make run"
          '';
        };
      });
}
