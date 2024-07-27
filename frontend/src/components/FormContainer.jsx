import {Row, Col} from "react-bootstrap"

const FormContainer = ({children})=> {
    return(
        <Row className="justify-content-center my-3">
            <Col xs={12} md={6} sm={4}>
                {children}
            </Col>
        </Row>
    )
}

export default FormContainer;