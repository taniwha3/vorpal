type artifact = { name : string; source : string }

let store = "/tmp/vorpal/store"
let example : artifact = { name = "example"; source = "." }

let create_directory store artifact_name =
  let dir_path = Filename.concat store artifact_name in
  Unix.mkdir dir_path 0o777;
  dir_path

let copy_file src dst =
  let ic = open_in src in
  let oc = open_out dst in
  try
    while true do
      output_char oc (input_char ic)
    done
  with End_of_file ->
    close_in ic;
    close_out oc

let rec copy_dir src dst ignore_files =
  if Sys.is_directory src then (
    if not (Sys.file_exists dst) then Unix.mkdir dst 0o777;
    Sys.readdir src
    |> Array.iter (fun file ->
           if not (List.mem file ignore_files) then
             copy_dir (Filename.concat src file) (Filename.concat dst file)
               ignore_files))
  else if not (List.mem (Filename.basename src) ignore_files) then (
    copy_file src dst;
    Printf.printf "Copied %s to %s\n" src dst)

let copy_files_to_directory artifact_path artifact_source =
  copy_dir artifact_source artifact_path
    [ ".git"; ".gitignore"; ".direnv"; "_build" ]

let () =
  copy_files_to_directory (create_directory store example.name) example.source
