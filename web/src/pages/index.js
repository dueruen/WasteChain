import React, { Fragment } from 'react';
import { Router } from '@reach/router';
import { PageContainer } from '../components';

import LandingPage from './landingPage'
import LoginPage from './login'
import CreateEmployeePage from './createAccount'
import CreateCompanyPage from './createCompany'
import CreateShipmentPage from './shipment/createShipment'
import StartTransferShipmentPage from './shipment/startTransferShipment'
import FinishTransferShipmentPage from './shipment/finishTransferShipment'
import DeatailedShipmentPage from './shipment/shipmentDetails'
import GetIDPage from './getID';

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <CreateEmployeePage path="createemployee" />
                    <CreateCompanyPage path="createcompany" />
                    <CreateShipmentPage path="createshipment" />
                    <StartTransferShipmentPage path="starttransfer" />
                    <FinishTransferShipmentPage path="finishtransfer" />
                    <DeatailedShipmentPage path="shipment/:shipmentID"/>
                    <GetIDPage path="getID" />
                </Router>
            </PageContainer>
        </Fragment>
    )
}
