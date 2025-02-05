// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
export function getColor(item: string) {
  if (
    item === "Successful" ||
    item === "Success" ||
    item === "Unused" ||
    item === "Enabled"
  )
    return "success";
  if (item === "Reserved") return "warning";
  else return "error";
}