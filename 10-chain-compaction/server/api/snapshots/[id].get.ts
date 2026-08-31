import { defineEventHandler } from "h3";
import { db } from "~/server/db";

export default defineEventHandler(async (event) => {
  const id = event.context.params?.id;
  if (!id) return { error: "missing id" };

  const row = await db.query(`SELECT * FROM snapshots WHERE snapshot_id = $1`, [id]);
  if (row.rowCount === 0) return { error: "snapshot_not_found" };

  return row.rows[0];
});
