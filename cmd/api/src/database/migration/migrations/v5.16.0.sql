-- Copyright 2024 Specter Ops, Inc.
--
-- Licensed under the Apache License, Version 2.0
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.
--
-- SPDX-License-Identifier: Apache-2.0


CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_saved_queries_description ON saved_queries using gin(description gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_saved_queries_name ON saved_queries USING gin(name gin_trgm_ops);

INSERT INTO parameters (id, key, name, description, value, created_at, updated_at) 
VALUES (3, 'analysis.citrix_rdp_support', 'Citrix RDP Support', 'This configuration parameter toggles Citrix support during post-processing. When on, CanRDP edges will come from the `Direct Access Users` group instead of the builtin `Remote Desktop Users` group.', '{"enabled": false}', current_timestamp, current_timestamp) ON CONFLICT DO NOTHING;