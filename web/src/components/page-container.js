import React, { Fragment } from 'react';
import styled from '@emotion/styled'
import { Navbar } from '.';

export default function PageContainer(props) {
    return (
        <Fragment>
            <Navbar/>
            <Container>{props.children}</Container>
        </Fragment>
    )
}

const Container =  styled.div`
    display: flex,
    align-items: center;
    justify-content: center;
    flex-direction: column;
    flexGrow: 1,
    width: 100%,
    maxWidth: 600,
    margin: 0 auto,
    padding: unit * 3,
    paddingBottom: unit * 5,
`
