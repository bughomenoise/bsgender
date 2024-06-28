{
  description = "bsgender";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  # inputs.go.url = "github:nixos/nixpkgs/a343533bccc62400e8a9560423486a3b6c11a23b";

  outputs = { nixpkgs, flake-utils, ... }@inputs:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];
          nativeBuildInputs = [
            # inputs.go.legacyPackages.${system}.go_1_22
          ];
          shellHook = ''
            go mod tidy
            go mod verify
          '';
        };

      });
}
