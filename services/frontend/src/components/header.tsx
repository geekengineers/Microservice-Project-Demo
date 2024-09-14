export function HeaderComponent() {
    return (
        <div className="navbar navbar-sticky navbar-glass">
            <div className="navbar-start">
                <a className="navbar-item">Microservice Project Demo</a>
            </div>
            <div className="navbar-end">
                <a className="navbar-item">Home</a>
                <a className="navbar-item">Blog</a>
                <a className="navbar-item">Products</a>
                <a className="btn btn-primary">Login</a>
            </div>
        </div>
    )
}