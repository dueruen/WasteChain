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
import EmployeeShipmentPage from './shipment/employeeShipmentPage'
import ProcessShipmentPage from './shipment/processShipment'
import NotFound from '../pages/notFound'

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <StartTransferShipmentPage path="shipment/transfer/:shipmentID" />
                    <FinishTransferShipmentPage path="shipment/finishtransfer" />
                    <GetIDPage path="my-id" />
                    <CreateEmployeePage path="employee/create" />
                    <EmployeeShipmentPage path="my-shipments/" />
                    <CreateCompanyPage path="company/create" />
                    <CreateShipmentPage path="shipment/create" />
                    <ProcessShipmentPage path="shipment/process/:shipmentID"/>
                    <DeatailedShipmentPage path="shipment/:shipmentID"/>
                    <NotFound default/>
                </Router>
            </PageContainer>
        </Fragment>
    )
}
