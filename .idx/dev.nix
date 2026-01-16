# Project IDX environment configuration
# https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  channel = "stable-23.11";

  packages = [
    pkgs.go
    pkgs.nodejs_20
    pkgs.nodePackages.nodemon
  ];

  env = {};

  idx = {
    extensions = [
      "golang.go"
      "google.gemini-cli-vscode-ide-companion"
    ];

    workspace = {
      onCreate = {
        open = {
          openFiles = [
            "core-example/cmd/server/main.go"
            "README.md"
          ];
        };
      };
    };

    previews = {
      enable = true;
    };
  };
}
