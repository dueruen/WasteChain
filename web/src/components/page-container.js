import React, { Fragment } from 'react';
import styled from '@emotion/core';
import { Navbar } from '.';

export default function PageContainer(props) {
    return (
        <Fragment>
            <Navbar/>
            <h1>Page container</h1>
        </Fragment>
    )
}
