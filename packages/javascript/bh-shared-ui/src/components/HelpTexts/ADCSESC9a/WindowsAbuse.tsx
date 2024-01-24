// Copyright 2024 Specter Ops, Inc.
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

import { FC } from 'react';
import makeStyles from '@mui/styles/makeStyles';
import { Typography, Link, List, ListItem, Box } from '@mui/material';

const useStyles = makeStyles((theme) => ({
    containsCodeEl: {
        '& code': {
            backgroundColor: 'darkgrey',
            padding: '2px .5ch',
            fontWeight: 'normal',
            fontSize: '.875em',
            borderRadius: '3px',
            display: 'inline',

            overflowWrap: 'break-word',
            whiteSpace: 'pre-wrap',
        },
    },
}));

const WindowsAbuse: FC = () => {
    const classes = useStyles();
    const step1 = (
        <>
            <Typography variant='body2' className={classes.containsCodeEl}>
                <b>Step 1: </b>Set UPN of victim to targeted principal's <code>sAMAccountName</code>.
                <br />
                <br />
                Set the UPN of the victim principal using PowerView:
            </Typography>
            <Typography component={'pre'}>
                {"Set-DomainObject -Identity VICTIM -Set @{'userprincipalname'='Target'}"}
            </Typography>
        </>
    );

    const step2 = (
        <>
            <Typography variant='body2' className={classes.containsCodeEl}>
                <b>Step 2: </b>Check if mail attribute of victim must be set and set it if required. If the certificate
                <br />
                <br />
                template is of schema version 2 or above and its attribute <code>msPKI-CertificateNameFlag</code>{' '}
                contains the flag <code>SUBJECT_REQUIRE_EMAIL</code> and/or <code>SUBJECT_ALT_REQUIRE_EMAIL</code> then
                the victim principal must have their mail attribute set for the certificate enrollment. The CertTemplate
                BloodHound node will have <em>"Subject Require Email"</em> or{' '}
                <em>"Subject Alternative Name Require Email"</em> set to true if any of the flags are present.
                <br />
                <br />
                If the certificate template is of schema version 1 or does not have any of the email flags, then
                continue to Step 3.
                <br />
                <br />
                If any of the two flags are present, you will need the victim’s mail attribute to be set. The value of
                the attribute will be included in the issues certificate but it is not used to identify the target
                principal why it can be set to any arbitrary string.
                <br />
                <br />
                Check if the victim has the mail attribute set using PowerView:
            </Typography>
            <Typography component={'pre'}>{'Get-DomainObject -Identity VICTIM -Properties mail'}</Typography>
            <Typography variant='body2'>
                If the victim has the mail attribute set, continue to Step 3.
                <br />
                <br />
                If the victim does not has the mail attribute set, set it to a dummy mail using PowerView:
            </Typography>
            <Typography component={'pre'}>
                {"Set-DomainObject -Identity VICTIM -Set @{'mail'='dummy@mail.com'}"}
            </Typography>
        </>
    );

    const step3 = (
        <Box
            sx={{
                borderRadius: '4px',
                backgroundColor: '#eee',
            }}>
            <Typography variant='body2' className={classes.containsCodeEl} sx={{ marginBottom: '-8px' }}>
                <b>Step 3: </b>Obtain a session as victim.
                <br />
                <br />
                There are several options for this step.
                <br />
                <br />
                If the victim is a computer, you can obtain the credentials of the computer account using the Shadow
                Credentials attack (see{' '}
                <Link
                    target='blank'
                    rel='noopener'
                    href='https://support.bloodhoundenterprise.io/hc/en-us/articles/17358104809499-AddKeyCredentialLink'>
                    AddKeyCredentialLink edge documentation
                </Link>
                ). Alternatively, you can obtain a session as SYSTEM on the host, which allows you to interact with AD
                as the computer account, by abusing control over the computer AD object (see{' '}
                <Link
                    target='blank'
                    rel='noopener'
                    href='https://support.bloodhoundenterprise.io/hc/en-us/articles/17312347318043-GenericAll'>
                    GenericAll edge documentation
                </Link>
                ).
                <br />
                <br />
                If the victim is a user, you have the following options for obtaining the credentials:
            </Typography>
            <List sx={{ fontSize: '12px' }}>
                <ListItem>
                    Shadow Credentials attack (see{' '}
                    <Link
                        target='blank'
                        rel='noopener'
                        href='https://support.bloodhoundenterprise.io/hc/en-us/articles/17358104809499-AddKeyCredentialLink'>
                        AddKeyCredentialLink edge documentation
                    </Link>
                    )
                </ListItem>
                <ListItem>
                    Password reset (see{' '}
                    <Link
                        target='blank'
                        rel='noopener'
                        href='https://support.bloodhoundenterprise.io/hc/en-us/articles/17223286750747-ForceChangePassword'>
                        ForceChangePassword edge documentation
                    </Link>
                    )
                </ListItem>
                <ListItem>
                    Targeted Kerberoasting (see{' '}
                    <Link
                        target='blank'
                        rel='noopener'
                        href='https://support.bloodhoundenterprise.io/hc/en-us/articles/17222775975195-WriteSPN'>
                        WriteSPN edge documentation
                    </Link>
                    )
                </ListItem>
            </List>
        </Box>
    );

    const step4 = (
        <>
            <Typography variant='body2'>
                <b>Step 4: </b>Enroll certificate as victim.
                <br />
                <br />
                Use Certify as the victim principal to request enrollment in the affected template, specifying the
                affected EnterpriseCA:
            </Typography>
            <Typography component={'pre'}>{'Certify.exe request /ca:SERVER\\CA-NAME /template:TEMPLATE'}</Typography>
            <Typography variant='body2' className={classes.containsCodeEl}>
                Save the certificate as <code>cert.pem</code> and the private key as <code>cert.key</code>.
            </Typography>
        </>
    );

    const step5 = (
        <>
            <Typography variant='body2'>
                <b>Step 5: </b>Convert the emitted certificate to PFX format:
            </Typography>
            <Typography component={'pre'}>{'certutil.exe -MergePFX .\\cert.pem .\\cert.pfx'}</Typography>
        </>
    );
    const step6 = (
        <>
            <Typography variant='body2'>
                <b>Step 6: </b>Set UPN of victim to arbitrary value.
                <br />
                <br />
                Set the UPN of the victim principal using PowerView:
            </Typography>
            <Typography component={'pre'}>
                {"Set-DomainObject -Identity VICTIM -Set @{'userprincipalname'='victim@corp.local'}"}
            </Typography>
        </>
    );
    const step7 = (
        <>
            <Typography variant='body2'>
                <b>Step 7: </b>Perform Kerberos authentication as targeted principal against affected DC using
                certificate.
                <br />
                <br />
                Use Rubeus to request a ticket granting ticket (TGT) from an affected DC, specifying the target identity
                to impersonate and the PFX-formatted certificate created in Step 5:
            </Typography>
            <Typography component={'pre'}>
                {'Rubeus.exe asktgt /certificate:cert.pfx /user:TARGET /domain:DOMAIN /dc:DOMAIN_CONTROLLER'}
            </Typography>
        </>
    );

    return (
        <>
            <Typography variant='body2'>An attacker may perform this attack in the following steps:</Typography>
            {step1}
            {step2}
            {step3}
            {step4}
            {step5}
            {step6}
            {step7}
        </>
    );
};

export default WindowsAbuse;
