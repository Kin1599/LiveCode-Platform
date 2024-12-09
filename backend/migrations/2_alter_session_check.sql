ALTER TABLE "Sessions"
DROP "Sessions_access_type_check";
ADD CHECK("access_type" IN("Public", "Private"));