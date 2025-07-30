# Final Project Documentation: Car Price Prediction

## Project Objective
The objective of this project is to build the most accurate machine learning model for predicting car prices based on various attributes. This process involves a systematic series of experiments to explore different techniques and find the best-performing model.

## Approach and Key Decisions
Throughout the project, several important decisions were made that directed the workflow:

### Evaluation Metrics Selection
**Primary Metric: Mean Absolute Error (MAE).** MAE was chosen as the primary metric because it is very easy to interpret. Its value represents the average prediction error in original units (Dollars). This provides a clear and direct picture of "how far" the model's predictions are from the actual prices. Our goal is to minimize this value.

**Supporting Metric: R2 Score (Coefficient of Determination).** R2 Score was chosen to provide context. This metric measures how much percentage of price variation can be explained by the model. Values close to 1.0 indicate a model that is very good at explaining the data.

### Initial Feature Selection
Initially, the car_ID and CarName columns were removed from the feature set.
- **Reason:** car_ID is a unique identifier (like a serial number) and has no predictive power. CarName has too many unique values, which if directly converted to dummy variables would create too many columns and make the model inefficient. This decision was later re-evaluated in the feature engineering stage.

## Experiment Summary and Results
Here are the steps that have been taken, along with findings and key metrics at each stage:

### Stage 1: Basic Model (Baseline)
**Description:** This initial stage focused on data preparation and training a basic model to establish a performance benchmark.

**Data Cleaning Process:**
- **Missing Values Check:** The first and most crucial cleaning step was checking for missing data. Using df.isnull().sum(), we found that this dataset is very clean and has no missing values in any column.
- **Irrelevant Column Removal:** car_ID and CarName columns were removed. car_ID only functions as an index, while CarName has too much variation to be used directly.
- **Data Type Conversion:** Categorical features (like fueltype, carbody, etc.) were converted to numeric format using pd.get_dummies() so they can be processed by machine learning models.

**Results:** After the data was clean and ready, several standard models were trained for comparison: Linear Regression, Ridge, Lasso, and Random Forest. Among these four models, Random Forest showed the best performance with a significant margin.

**Key Metrics (Random Forest):** MAE ~$1261 | R2 Score: ~0.958

### Stage 2: Improvement 1 - Feature Engineering (Adding Car Brand)
**Description:** Extracting car brand information from the CarName column to create a new brand feature.

**Results:** Linear model performance improved dramatically, proving that brand is an important feature. Random Forest performance did not change much.

**Key Metrics (Random Forest):** MAE ~$1288 | R2 Score: ~0.958

### Stage 3: Improvement 2 - Random Forest Hyperparameter Optimization
**Description:** Attempting to refine the Random Forest model from Stage 2 using RandomizedSearchCV.

**Results:** The tuned model actually showed worse performance (MAE ~$1472). It was concluded that default settings are more general.

### Stage 4: Improvement 3 - Logarithmic Transformation on RF
**Description:** Applying logarithmic transformation to the price variable to address skewed data distribution.

**Results:** This technique also failed to improve Random Forest performance (MAE ~$1356), indicating the model is already robust enough.

### Stage 5: Improvement 4 - XGBoost Experiment
**Description:** Trying the popular gradient boosting model, XGBoost, with default settings.

**Results:** XGBoost was unable to beat Random Forest performance (MAE ~$1520).

### Stage 6: Final Experiment (Ridge & Lasso Tuned + Log)
**Description:** Giving the best chance to linear models by tuning alpha and applying log transformation to the target.

**Results:** Performance improved from the basic version, but still far below Random Forest (MAE Ridge ~$1475, MAE Lasso ~$1571).

### Stage 7: Ultimate Experiment (XGBoost Tuned + Log)
**Description:** Combining all techniques (XGBoost, tuning, log transformation) to create the strongest challenger.

**Results:** Even when fully optimized, XGBoost still couldn't beat Random Forest (MAE ~$1588).

### Stage 8 (Deployment): Model Conversion to ONNX
**Description:** As preparation for deployment in non-Python environments (like Go), the winning scikit-learn model (.joblib) was converted to the interoperable ONNX format.

**Process:** Using the skl2onnx library, the model was converted by defining appropriate input signatures.

**Results:** Conversion was successful, and verification was performed to ensure the ONNX model provides identical predictions to the original model. The final deployment-ready artifact is best_model.onnx.

## Final Conclusion and Winning Model
After going through a systematic series of experiments, we concluded that the best model for predicting car prices on this dataset is:

🏆 **Winning Model: Random Forest Regressor with additional brand feature (from Stage 2).**

**Final MAE:** $1288.82  
**Final R2 Score:** 0.958

### Why the Model with Brand Feature was Chosen (MAE ~1288 vs. Initial Model MAE ~1261)?
Although the initial model had a slightly lower MAE value, the model with brand feature was chosen as the winner for several strategic reasons that are more important than a small numerical difference:

1. **Insignificant Difference:** The ~$27 difference is most likely statistical "noise" and could reverse if the data were split differently.
2. **More Logical and Robust Model:** The winning model explicitly understands the concept of "car brand", which is a major price driver in the real world. This makes it more reliable and easier to interpret.
3. **Better Generalization:** Models built on more logical feature foundations tend to have better generalization capabilities on new data in the future.