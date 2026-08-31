import { defineEventHandler, readBody } from "h3";
import { spawn } from "node:child_process";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const { segment_start, segment_end } = body;

  if (!segment_start || !segment_end) {
    return { error: "segment_start and segment_end required" };
  }

  const proc = spawn("./nitro-compact", ["--start", segment_start, "--end", segment_end], {
    stdio: ["ignore", "pipe", "pipe"],
  });

  let out = "";
  let errOut = "";
  for await (const chunk of proc.stdout) out += chunk.toString();
  for await (const chunk of proc.stderr) errOut += chunk.toString();

  if (errOut) {
    return { error: errOut };
  }

  return { result: out.trim() };
});
