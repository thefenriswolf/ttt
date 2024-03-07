{ stdenvNoCC, lib, fetchFromGitHub, makeWrapper, buildGoModule }:

buildGoModule rec {
  pname = "ttt";
  version = "v20240307";

  src = fetchFromGitHub {
    owner = "thefenriswolf";
    repo = "ttt";
    rev = "${version}";
    hash = "sha256-d3oyPLh9vPNmbbT1qorc59a1cT3KFR89cyI67XiXsHU=";
  };
  vendorHash = "sha256-ekZ5rRbvD8U+UEfqWbPCZ9v++ZDTpAuU3LT9hWlwC5Q=";

  meta = with lib; {
    description = "Time Tracker Tool written in Golang";
    homepage = "https://github.com/thefenriswolf/ttt";
    license = licenses.bsd3;
    maintainers = with maintainers; [ thefenriswolf ];
  };
}
