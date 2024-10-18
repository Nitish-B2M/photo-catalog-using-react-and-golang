import "./App.css";
import Navbar from "./components/Navbar/Navbar";
import CatalogSection from "./components/Main/CatalogSection";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import GallerySection from "./components/Main/GallerySection";
import Authentication from "./components/Auth/Authentication";
import ProfilePage from "./components/Profile/ProfilePage";
import PageNotFound from "./components/Error/PageNotFound";
import PrivateRoute from "./components/Auth/PrivateRoute";

function App() {
  const authPaths = ["/auth", "/login", "/register"];

  return (
    <>
      <Router>
        <div className="font-primary overflow-x-hidden">
          <Navbar />
          <Routes>
            <Route
              path="/"
              element={
                <>
                  <CatalogSection />
                </>
              }
            />
            <Route
              path="/gallery"
              element={
                <PrivateRoute>
                  <GallerySection />
                </PrivateRoute>
              }
            />
            {authPaths.map((path) => (
              <Route key={path} path={path} element={
                <>
                  <Authentication currentPath={path}/>
                </>
              } />
            ))}
            <Route
              path="/profile"
              element={
                <PrivateRoute>
                  <ProfilePage />
                </PrivateRoute>
              }
            />
            <Route
              path="*"
              element={
                <>
                  <PageNotFound />
                </>
              }
            />
          </Routes>
        </div>
      </Router>
    </>
  );
}

export default App;
